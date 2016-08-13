/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package services

import (
	"github.ibm.com/riethm/gopherlayer/datatypes"
	"github.ibm.com/riethm/gopherlayer/session"
	"github.ibm.com/riethm/gopherlayer/sl"
)

//
type Marketplace_Partner struct {
	Session *session.Session
	Options sl.Options
}

func GetMarketplacePartnerService(sess *session.Session) *Marketplace_Partner {
	return &Marketplace_Partner{Session: sess}
}

func (r Marketplace_Partner) Id(id int) *Marketplace_Partner {
	r.Options.Id = &id
	return &r
}

func (r Marketplace_Partner) Mask(mask string) *Marketplace_Partner {
	r.Options.Mask = mask
	return &r
}

func (r Marketplace_Partner) Filter(filter string) *Marketplace_Partner {
	r.Options.Filter = filter
	return &r
}

func (r Marketplace_Partner) Limit(limit int) *Marketplace_Partner {
	r.Options.Limit = &limit
	return &r
}

func (r Marketplace_Partner) Offset(offset int) *Marketplace_Partner {
	r.Options.Offset = &offset
	return &r
}

//
func (r *Marketplace_Partner) GetAllObjects() (resp []datatypes.Marketplace_Partner, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Marketplace_Partner) GetAllPublishedPartners(searchTerm *string) (resp []datatypes.Marketplace_Partner, err error) {
	params := []interface{}{
		searchTerm,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Marketplace_Partner) GetAttachments() (resp []datatypes.Marketplace_Partner_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Marketplace_Partner) GetFeaturedPartners(non *bool) (resp []datatypes.Marketplace_Partner, err error) {
	params := []interface{}{
		non,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

//
func (r *Marketplace_Partner) GetFile(name *string) (resp datatypes.Marketplace_Partner_File, err error) {
	params := []interface{}{
		name,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Marketplace_Partner) GetLogoMedium() (resp datatypes.Marketplace_Partner_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Marketplace_Partner) GetLogoMediumTemp() (resp datatypes.Marketplace_Partner_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Marketplace_Partner) GetLogoSmall() (resp datatypes.Marketplace_Partner_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Marketplace_Partner) GetLogoSmallTemp() (resp datatypes.Marketplace_Partner_Attachment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Marketplace_Partner) GetObject() (resp datatypes.Marketplace_Partner, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Marketplace_Partner) GetPartnerByUrlIdentifier(urlIdentifier *string) (resp datatypes.Marketplace_Partner, err error) {
	params := []interface{}{
		urlIdentifier,
	}
	err = invokeMethod(params, r.Session, &r.Options, &resp)
	return
}