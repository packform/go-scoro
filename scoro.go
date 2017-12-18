// Package scoro provides Go client code for sending requests to Scoro API.
//
// Services
//
// Code of the library is divided into several services, one for each
// Scoro API "module". Each service provides a subset of
// View/List/Modify/Delete actions which are directly mapped to the
// corresponding API calls.
//
// Products service:
//
//    products := scoro.Products(credentials)
//
// Quotes service:
//
//    quotes := scoro.Quotes(credentials)
//
// Orders service:
//
//    orders := scoro.Orders(credentials)
//
// Invoices service:
//
//    invoices := scoro.Invoices(credentials)
//
package scoro
