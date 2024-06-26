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

output "ids" {
  value       = module.public_dns_zone.ids
  description = "The DNS Zone IDs."
}

output "max_number_of_record_sets" {
  value       = module.public_dns_zone.max_number_of_record_sets
  description = "(Optional) Maximum number of Records in the zone. Defaults to 1000."
}

output "number_of_record_sets" {
  value       = module.public_dns_zone.number_of_record_sets
  description = "(Optional) The number of records already in the zone."
}

output "name_servers" {
  value       = module.public_dns_zone.name_servers
  description = "(Optional) A list of values that make up the NS record for the zone."
}

output "resource_group_name" {
  value       = module.resource_group.name
  description = "(Optional) The name of the resource group in which the zone exists."
}
