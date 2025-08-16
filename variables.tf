# variables.tf

variable "prefix" {
  type        = list(string)
  default     = []
  description = "It is not recommended that you use prefix by azure you should be using a suffix for your resources."
}

variable "suffix" {
  type        = list(string)
  default     = []
  description = "It is recommended that you specify a suffix for consistency. please use only lowercase characters when possible"
}

variable "unique-seed" {
  description = "Custom value for the random characters to be used"
  type        = string
  default     = ""
}

variable "unique-length" {
  description = "Max length of the uniqueness suffix to be added"
  type        = number
  default     = 4
}

variable "unique-include-numbers" {
  description = "If you want to include numbers in the unique generation"
  type        = bool
  default     = true
}

// --- New Variables for Environments and Locations ---

variable "environment" {
  description = "The environment for the resource (e.g., 'dev', 'prod'). This is for documentation and for use in name components."
  type        = string
  default     = ""
}

variable "location" {
  description = "The location for the resource. Can be the full name ('eastus'), short name ('eus'), or display name ('East US'). The module will look up the correct value."
  type        = string
  default     = ""
}