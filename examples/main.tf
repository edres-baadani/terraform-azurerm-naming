
// Example 1

module "name_empty" {
  source = "../"
}

output "name_empty" {
  value = module.name_empty.storage_account.name_unique
}

// Example 2

module "suffix" {
  source        = "../"
  suffix        = ["su", "fix"]
  unique-length = 20
}

output "suffix" {
  value = module.suffix.storage_account.name_unique
}

// Example 3

module "random" {
  source      = "../"
  unique-seed = module.suffix.unique-seed
}

output "random" {
  value = module.random.storage_account.name_unique
}

// Example 4

module "everything" {
  source                 = "../"
  prefix                 = ["pre", "fix"]
  suffix                 = ["su", "fix"]
  unique-seed            = "random"
  unique-length          = 2
  unique-include-numbers = false
}

output "everything" {
  value = module.everything.storage_account.name_unique
}

output "validation_everything" {
  value = module.everything.validation
}

// Example 5: Using Environments Data

variable "environment" {
  description = "The environment for the resource."
  type        = string
  default     = "prod" // Can be "dev", "prod", "stg", etc.
}

// Pass the environment variable to the naming module
module "with_environment" {
  source      = "../"
  prefix      = ["myorg", var.environment]
  suffix      = ["myapp"]
}

// Get the environment details from the module's outputs
output "environments_lookup" {
  description = "Example of looking up environment details."
  value       = module.with_environment.environments[var.environment]
}

// Show the full name with the environment included
output "with_environment_name" {
  description = "Example of a resource name with an environment prefix."
  value       = module.with_environment.storage_account.name_unique
}

// Example 6: Using Locations Data

variable "location" {
  description = "The location for the resource. Can be name, short name, or display name."
  type        = string
  default     = "eastus" // Can be "eus", "East US", etc.
}

// Pass the location variable to the naming module
module "with_location" {
  source      = "../"
  prefix      = ["myorg"]
  suffix      = ["myapp"]
  location    = var.location
}

// Show the details of the looked-up location
output "locations_lookup" {
  description = "Example of looking up location details."
  value       = module.with_location.locations[var.location]
}

// Show the full name with the location included
output "with_location_name" {
  description = "Example of a resource name with a location suffix."
  value       = module.with_location.storage_account.name_unique
}