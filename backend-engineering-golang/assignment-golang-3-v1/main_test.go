package main_test

import (
	main "a21hc3NpZ25tZW50"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RecordInvoice", func() {
	// =====================================================================
	// FinanceInvoice RecordInvoice test data
	// =====================================================================
	Describe("FinanceInvoice", func() {
		When("date is empty string", func() {
			It("should return error 'invoice date is empty'", func() {
				financeInvoice := main.FinanceInvoice{
					Date:     "",
					Status:   main.PAID,
					Approved: true,
					Details:  []main.Detail{},
				}
				_, err := financeInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice date is empty")))
			})
		})
		When("status is empty or not 'paid' or not 'unpaid'", func() {
			It("should return error 'invoice status is not valid'", func() {
				financeInvoice := main.FinanceInvoice{
					Date:     "01/01/2022",
					Status:   "",
					Approved: true,
					Details:  []main.Detail{},
				}
				_, err := financeInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice status is not valid")))
			})
		})
		When("details is empty data", func() {
			It("should return error 'invoice details is empty'", func() {
				financeInvoice := main.FinanceInvoice{
					Date:     "01/01/2022",
					Status:   main.PAID,
					Approved: true,
					Details:  []main.Detail{},
				}
				_, err := financeInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice details is empty")))
			})
		})
		When("total in details is 0 or less than 0", func() {
			It("should return error 'total price is not valid'", func() {
				financeInvoice := main.FinanceInvoice{
					Date:     "01/01/2022",
					Status:   main.PAID,
					Approved: true,
					Details:  []main.Detail{{"pembelian nota", 0}, {"Pembelian alat tulis", -200}},
				}
				_, err := financeInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("total price is not valid")))
			})
		})
		When("invoice data as needed", func() {
			It("should return return error and covert invoice to InvoiceData", func() {
				financeInvoice := main.FinanceInvoice{
					Date:     "01/01/2022",
					Status:   main.PAID,
					Approved: true,
					Details:  []main.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
				}
				inv, err := financeInvoice.RecordInvoice()
				Expect(err).To(BeNil())
				Expect(inv).To(Equal(main.InvoiceData{
					Date:         "01-January-2022",
					TotalInvoice: 8000,
					Departemen:   main.Finance,
				}))
			})
		})
	})

	// =====================================================================
	// WarehouseInvoice RecordInvoice test data
	// =====================================================================
	Describe("WarehouseInvoice", func() {
		When("date is empty string", func() {
			It("should return error 'invoice date is empty'", func() {
				warehouseInvoice := main.WarehouseInvoice{
					Date:        "",
					InvoiceType: main.PURCHASE,
					Approved:    true,
					Products: []main.Product{
						{"product A", 10, 10000, 0.1},
						{"product C", 5, 15000, 0.2},
					},
				}
				_, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice date is empty")))
			})
		})

		When("invoice type is empty string or not 'pruchase' or not 'sales'", func() {
			It("should return error 'invoice type is not valid'", func() {
				warehouseInvoice := main.WarehouseInvoice{
					Date:        "01-January-2022",
					InvoiceType: "",
					Approved:    true,
					Products: []main.Product{
						{"product A", 10, 10000, 0.1},
						{"product C", 5, 15000, 0.2},
					},
				}
				_, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice type is not valid")))
			})
		})

		When("invoice product is empty data", func() {
			It("should return error 'invoice products is empty'", func() {
				warehouseInvoice := main.WarehouseInvoice{
					Date:        "01-January-2022",
					InvoiceType: "purchase",
					Approved:    true,
					Products:    []main.Product{},
				}
				_, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice products is empty")))
			})
		})

		When("invoice product unit is 0 or lower than 0", func() {
			It("should return error 'unit product is not valid'", func() {
				warehouseInvoice := main.WarehouseInvoice{
					Date:        "01-January-2022",
					InvoiceType: "purchase",
					Approved:    true,
					Products: []main.Product{
						{"product A", -1, 10000, 0.1},
						{"product C", 0, 15000, 0.2},
					},
				}
				_, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("unit product is not valid")))
			})
		})

		When("invoice price product is 0 or less than 0", func() {
			It("should return error 'price product is not valid'", func() {
				warehouseInvoice := main.WarehouseInvoice{
					Date:        "01-January-2022",
					InvoiceType: "purchase",
					Approved:    true,
					Products: []main.Product{
						{"product A", 10, -1, 0.1},
						{"product C", 5, 0, 0.2},
					},
				}
				_, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("price product is not valid")))
			})
		})

		When("invoice warehouse data as needed", func() {
			It("should not return error and convert data to Invoice Data", func() {
				warehouseInvoice := main.WarehouseInvoice{
					Date:        "01-January-2022",
					InvoiceType: "purchase",
					Approved:    true,
					Products: []main.Product{
						{"product A", 10, 10000, 0.1},
						{"product C", 5, 15000, 0.2},
					},
				}
				inv, err := warehouseInvoice.RecordInvoice()
				Expect(err).To(BeNil())
				Expect(inv).To(Equal(main.InvoiceData{
					Date:         "01-January-2022",
					TotalInvoice: 150000,
					Departemen:   main.Warehouse,
				}))
			})
		})
	})

	// =====================================================================
	// MarketingInvoice RecordInvoice test data
	// =====================================================================
	Describe("MarketingInvoice", func() {
		When("date is empty string", func() {
			It("should return error 'invoice date is empty'", func() {
				MarketingInvoice := main.MarketingInvoice{
					Date:     "",
					Approved: true,
				}
				_, err := MarketingInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("invoice date is empty")))
			})
		})

		When("start date and end date is emptu", func() {
			It("should return error 'travel date is empty'", func() {
				MarketingInvoice := main.MarketingInvoice{
					Date:      "01/01/2022",
					StartDate: "",
					EndDate:   "",
					Approved:  true,
				}
				_, err := MarketingInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("travel date is empty")))
			})
		})

		When("PricePerDay is 0 or less than 0", func() {
			It("should return error 'price per day is not valid'", func() {
				MarketingInvoice := main.MarketingInvoice{
					Date:        "01/01/2022",
					StartDate:   "01/01/2022",
					EndDate:     "02/01/2022",
					PricePerDay: -1000,
					Approved:    true,
				}
				_, err := MarketingInvoice.RecordInvoice()
				Expect(err).To(Equal(errors.New("price per day is not valid")))
			})
		})

		When("invoice marketing data as needed", func() {
			It("should not return error and convert data to Invoice Data", func() {
				MarketingInvoice := main.MarketingInvoice{
					Date:        "03/01/2022",
					StartDate:   "01/01/2022",
					EndDate:     "02/01/2022",
					PricePerDay: 10000,
					AnotherFee:  5000,
					Approved:    true,
				}
				inv, err := MarketingInvoice.RecordInvoice()
				Expect(err).To(BeNil())
				Expect(inv).To(Equal(main.InvoiceData{
					Date:         "03-January-2022",
					TotalInvoice: 25000,
					Departemen:   main.Marketing,
				}))
			})
		})
	})
})

var _ = Describe("RecapDataInvoice", func() {
	// test same date same departemen
	When("all invoice data is same date and same departemen", func() {
		It("should not return error, and return one data invoice", func() {
			invoices := []main.Invoice{
				main.FinanceInvoice{
					Date:     "01/02/2020",
					Details:  []main.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
					Status:   main.PAID,
					Approved: true,
				},
				main.FinanceInvoice{
					Date:     "01/02/2020",
					Details:  []main.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
					Status:   main.PAID,
					Approved: true,
				},
			}

			listInv, err := main.RecapDataInvoice(invoices)
			Expect(err).To(BeNil())

			Expect(listInv).To(Equal([]main.InvoiceData{
				{"01-February-2020", 16000, main.Finance},
			}))
		})
	})
	// test same date different departemen
	When("all invoice data is same date and difference departemen", func() {
		It("should not return error, and return list data by departement", func() {
			invoices := []main.Invoice{
				main.FinanceInvoice{
					Date:     "01/02/2020",
					Details:  []main.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
					Status:   main.PAID,
					Approved: true,
				},
				main.FinanceInvoice{
					Date:     "01/02/2020",
					Details:  []main.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
					Status:   main.PAID,
					Approved: true,
				},
				main.WarehouseInvoice{
					Date: "01-February-2020",
					Products: []main.Product{
						{"product A", 10, 10000, 0.1},
						{"product C", 5, 15000, 0.2},
					},
					InvoiceType: main.PURCHASE,
					Approved:    true,
				},
				main.MarketingInvoice{
					Date:        "01/02/2020",
					StartDate:   "20/01/2020",
					EndDate:     "25/01/2020",
					Approved:    true,
					PricePerDay: 10000,
					AnotherFee:  5000,
				},
			}

			listInv, err := main.RecapDataInvoice(invoices)
			Expect(err).To(BeNil())

			expected := []main.InvoiceData{
				{"01-February-2020", 16000, main.Finance},
				{"01-February-2020", 65000, main.Marketing},
				{"01-February-2020", 150000, main.Warehouse},
			}

			for _, inv := range listInv {
				Expect(expected).To(ContainElement(inv))
			}
		})
	})

	// test different date same departemen
	When("all invoice data with differece date and difference departemen", func() {
		It("should not return error, and return list data by date and by departement", func() {
			invoices := []main.Invoice{
				main.FinanceInvoice{
					Date:     "01/02/2020",
					Details:  []main.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
					Status:   main.PAID,
					Approved: true,
				},
				main.FinanceInvoice{
					Date:     "02/02/2020",
					Details:  []main.Detail{{"pembelian nota", 4000}, {"Pembelian alat tulis", 4000}},
					Status:   main.PAID,
					Approved: true,
				},
				main.WarehouseInvoice{
					Date: "03-February-2020",
					Products: []main.Product{
						{"product A", 10, 10000, 0.1},
						{"product C", 5, 15000, 0.2},
					},
					InvoiceType: main.PURCHASE,
					Approved:    true,
				},
				main.MarketingInvoice{
					Date:        "04/02/2020",
					StartDate:   "20/01/2020",
					EndDate:     "25/01/2020",
					Approved:    true,
					PricePerDay: 10000,
					AnotherFee:  5000,
				},
			}

			listInv, err := main.RecapDataInvoice(invoices)
			Expect(err).To(BeNil())

			expected := []main.InvoiceData{
				{"01-February-2020", 8000, main.Finance},
				{"02-February-2020", 8000, main.Finance},
				{"03-February-2020", 150000, main.Warehouse},
				{"04-February-2020", 65000, main.Marketing},
			}

			for _, inv := range listInv {
				Expect(expected).To(ContainElement(inv))
			}
		})
	})
})
