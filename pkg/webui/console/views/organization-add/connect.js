// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { push } from 'connected-react-router'

import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'

import { createOrganization } from '@console/store/actions/organizations'

const mapDispatchToProps = dispatch => ({
  ...bindActionCreators(
    {
      createOrganization: attachPromise(createOrganization),
    },
    dispatch,
  ),
  createOrganizationSuccess: id => dispatch(push(`/organizations/${id}`)),
})

export default OrganizationAdd => connect(null, mapDispatchToProps)(OrganizationAdd)
