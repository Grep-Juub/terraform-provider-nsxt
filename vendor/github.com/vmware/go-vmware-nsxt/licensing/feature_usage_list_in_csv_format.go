/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package licensing

type FeatureUsageListInCsvFormat struct {

	// File name set by HTTP server if API  returns CSV result as a file.
	FileName string `json:"file_name,omitempty"`

	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`

	Results []FeatureUsageCsvRecord `json:"results,omitempty"`
}
