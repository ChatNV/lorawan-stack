// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React from 'react'
import bind from 'autobind-decorator'
import classnames from 'classnames'
import { getIn } from 'formik'

import Icon from '@ttn-lw/components/icon'

import Message from '@ttn-lw/lib/components/message'

import from from '@ttn-lw/lib/from'
import PropTypes from '@ttn-lw/lib/prop-types'

import FormContext from '../context'

import style from './field.styl'

export function getPassThroughProps(props, excludeProps) {
  const rest = {}
  for (const property of Object.keys(props)) {
    if (!excludeProps[property]) {
      rest[property] = props[property]
    }
  }
  return rest
}

const isValueEmpty = function(value) {
  if (value === null || value === undefined) {
    return true
  }

  if (typeof value === 'object') {
    return Object.keys(value) === 0
  }

  if (typeof value === 'string') {
    return value === ''
  }

  return false
}

@bind
class FormField extends React.Component {
  static contextType = FormContext

  static propTypes = {
    className: PropTypes.string,
    component: PropTypes.oneOfType([
      PropTypes.func,
      PropTypes.string,
      PropTypes.shape({
        render: PropTypes.func.isRequired,
      }),
    ]).isRequired,
    description: PropTypes.message,
    disabled: PropTypes.bool,
    name: PropTypes.string.isRequired,
    onChange: PropTypes.func,
    readOnly: PropTypes.bool,
    required: PropTypes.bool,
    title: PropTypes.message.isRequired,
    warning: PropTypes.message,
  }

  static defaultProps = {
    className: undefined,
    disabled: false,
    onChange: () => null,
    warning: '',
    description: '',
    readOnly: false,
    required: false,
  }

  componentDidMount() {
    const { name } = this.props

    this.context.registerField(name, this)
  }

  componentWillUnmount() {
    const { name } = this.props

    this.context.unregisterField(name)
  }

  extractValue(value) {
    let newValue = value
    if (typeof value === 'object' && 'target' in value) {
      const target = value.target
      if ('type' in target && target.type === 'checkbox') {
        newValue = target.checked
      } else if ('value' in target) {
        newValue = target.value
      }
    }

    return newValue
  }

  handleChange(value, enforceValidation = false) {
    const { name, onChange } = this.props
    const { setFieldValue, setFieldTouched } = this.context

    // Check if the value is react's synthetic event.
    const newValue = this.extractValue(value)

    setFieldValue(name, newValue)

    if (enforceValidation) {
      setFieldTouched(name)
    }

    onChange(value)
  }

  handleBlur(event) {
    const { name } = this.props
    const { validateOnBlur, setFieldTouched } = this.context

    if (validateOnBlur) {
      const value = this.extractValue(event)
      setFieldTouched(name, !isValueEmpty(value))
    }
  }

  render() {
    const {
      className,
      name,
      title,
      warning,
      description,
      disabled,
      required,
      readOnly,
      component: Component,
    } = this.props
    const { horizontal, disabled: formDisabled } = this.context

    const fieldValue = getIn(this.context.values, name)
    const fieldError = getIn(this.context.errors, name)
    const fieldTouched = getIn(this.context.touched, name)
    const fieldDisabled = disabled || formDisabled

    const hasError = Boolean(fieldError)
    const hasWarning = Boolean(warning)
    const hasDescription = Boolean(description)

    const showError = fieldTouched && hasError
    const showWarning = !hasError && hasWarning
    const showDescription = !showError && !showWarning && hasDescription

    const fieldMessage = showError ? (
      <div className={style.messages}>
        <Err content={fieldError} title={name} />
      </div>
    ) : showWarning ? (
      <div className={style.messages}>
        <Err content={warning} title={name} warning />
      </div>
    ) : showDescription ? (
      <Message className={style.description} content={description} />
    ) : null

    const fieldComponentProps = {
      value: fieldValue,
      error: showError,
      warning: showWarning,
      name,
      horizontal,
      disabled: fieldDisabled,
      onChange: this.handleChange,
      onBlur: this.handleBlur,
    }

    const cls = classnames(
      className,
      style.field,
      from(style, {
        error: showError,
        warning: showWarning,
        horizontal,
        required,
        readOnly,
      }),
    )

    return (
      <div className={cls} data-needs-focus={showError}>
        <label className={style.label}>
          <Message content={title} className={style.title} />
          <span className={style.reqicon}>&middot;</span>
        </label>
        <div className={style.componentArea}>
          <Component
            {...fieldComponentProps}
            {...getPassThroughProps(this.props, FormField.propTypes)}
          />
          {fieldMessage}
        </div>
      </div>
    )
  }
}

const Err = ({ content, error, warning, title, className }) => {
  const icon = error ? 'error' : 'warning'
  const contentValues = content.values || {}
  const classname = classnames(style.message, className, {
    [style.show]: content && content !== '',
    [style.hide]: !content || content === '',
    [style.err]: error,
    [style.warn]: warning,
  })

  if (title) {
    contentValues.field = <Message content={title} className={style.name} />
  }

  return (
    <div className={classname}>
      <Icon icon={icon} className={style.icon} />
      <Message
        content={content.format || content.error_description || content.message || content}
        values={contentValues}
      />
    </div>
  )
}

Err.propTypes = {
  className: PropTypes.string,
  content: PropTypes.error.isRequired,
  error: PropTypes.bool,
  title: PropTypes.message,
  warning: PropTypes.bool,
}

Err.defaultProps = {
  className: undefined,
  title: undefined,
  warning: false,
  error: true,
}

export default FormField
