

variable "reservation_id" {
  type    = string
  default = "nVDVdrB7YCrtGzwxuXjzyh"
}

# Returns all reservations in au/AustraliaEast
data "xtratusipam_reservations" "all" {
  space = "au"
  block = "AustraliaEast"
}
output "all_reservations" {
  value = data.xtratusipam_reservations.all.reservations
}

resource "xtratusipam_reservation" "created" {
  space       = "au"
  block       = "AustraliaEast"
  size        = 24
  description = "prueba fnieto"
}
output "created" {
  value = xtratusipam_reservation.created
}