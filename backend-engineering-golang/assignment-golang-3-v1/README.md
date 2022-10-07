# Assignment 3

## Invoice Reporting

### Materi yang diuji

- Struct
- Interface
- Error Handling

### Description

Terdapat perusahaan "kitaSemangat" yang memiliki banyak departemen dan bergerak di penjualan produk (retail). Perusahaan ini memiliki kendala dalam melakukan pelaporan faktur / _invoice_ setiap akhir bulan karena ada beberapa format _faktur_ yang berbeda-beda di setiap departemen. Khususnya terdapat 3 departement yang memiliki format faktur yang sangat berbeda, yaitu Finance, Warehouse, dan Marketing.

Buatlah program yang dapat melakukan rekap menjadi format invoice yang sudah disepakati bersama. Karena kita sudah belajar `struct`, maka format invoice dari tiap departemen akan di konversi menjadi sebagai berikut:

1. Finance

    Invoice finance umumnya berisi pembelian barang-barang. Terdapat data berupa tanggal, status, approve status dan Detail pembelian.

    - tanggal (Date) bertipe `string` berformat `"DD/MM/YYYY"` (contoh: "01/01/2020")
    - status bertipe `string` adalah informasi khusus di departemen keuangan yang menunjukkan apakah invoice tersebut sudah terbayar atau tidak
    - Approve status bertipe `bool` yang akan berisi informasi yang menandakan invoice tersebut sudah ditandatangani oleh atasan, jika bernilai `true` maka invoice sudah ditandatangani, dan akan `false` jika sebaliknya
    - Detail pembelian akan berisi deskripsi (bertipe `string`) dan total pembelian (bertipe `int`), di setiap invoice data ini bisa lebih dari 1

    ```go
    type FinanceInvoice struct {
        Date       string
        Status     InvoiceStatus // status: "paid", "unpaid"
        Approved   bool
        Details   []Detail
    }

    type InvoiceStatus string

    const (
        PAID   InvoiceStatus = "paid"
        UNPAID InvoiceStatus = "unpaid"
    )

    type Detail struct {
        Description string
        Total       int
    }
    ```

2. Warehouse

    Invoice departemen warehouse umumnya digunakan untuk mencatat barang masuk dan keluar termasuk harga dan diskonnya. Terdapat data berupa tanggal, tipe invoice, approve status dan Detail produk.

    - tanggal (Date) bertipe `string` berformat `"DD-month-YYYY"` (contoh: "01-january-2022")
    - tipe invoice bertipe `InvoiceTypeName` yang akan berisi informasi apakah invoice tersebut merupakan invoice pembelian atau penjualan, jika bernilai `purchase` maka invoice tersebut merupakan invoice pembelian, dan akan `sales` jika sebaliknya
    - Approve status bertipe `bool` yang sama seperti invoice dari departemen Finance
    - Detail product (Products) bertipe _struct_ yang akan berisi informasi product (nama, total, harga per product, dan diskon yang diberikan)

    ```go
    type WarehouseInvoice struct {
        Date     string
        InvoiceType InvoiceTypeName
        Approved bool
        Products []Product
    }

    type InvoiceTypeName string

    const (
        PURCHASE InvoiceTypeName = "purchase"
        SALES    InvoiceTypeName = "sales"
    )


    type Product struct {
        Name     string
        Unit    int
        Price    int
        Discount int
    }
    ```

3. Marketing

    Invoice marketing umumnya digunakan untuk keluar kota. Terdapat data berupa tanggal invoice, tanggal mulai perjalanan sampai selesai (untuk kebutuhan biaya penginapan), dan biaya lain yang harus dihitung juga, dan approve status.

    - Tanggal (Date) bertipe `string` berformat `"DD/MM/YYYY"` (contoh: "01/01/2020")
    - Tanggal mulai perjalanan (StartDate) bertipe `string` berformat `"DD/MM/YYYY"` (contoh: "01/01/2020")
    - Tanggal selesai perjalanan (EndDate) bertipe `string` berformat `"DD/MM/YYYY"` (contoh: "01/01/2020")
    - Biaya per hari (PricePerDay) bertipe `int` yang akan berisi informasi biaya penginapan per hari berdasarkan jarak tanggal mulai perjalanan sampai selesai
    - Biaya lain (AnotherFee) bertipe `int` yang akan berisi informasi biaya lain yang harus dibayarkan
    - Approve status bertipe `bool` yang sama seperti invoice dari departemen Finance dan Warehouse

    ```go
    type MarketingInvoice struct {
        Date        string
        StartDate   string
        EndDate     string
        PricePerDay int
        AnotherFee  int
        Approved    bool
    }
    ```

### Phase 1

Kita akan mengimplementasikan _interface_ untuk masing-masing _struct_ yang mewakili format invoice dari tiap departement.

```go
type Invoice interface {
    RecordInvoice() (InvoiceData, error)
}
```

Buatlah _method_ yang dibutuhkan untuk _Struct_ dari departemen finance agar dapat mengimplementasikan _interface_ `Invoice` di atas. Kembalikan `InvoiceData` dengan format sebagai berikut:

```go
type InvoiceData struct {
    Date         string
    TotalInvoice float64
    departemen   DepartmentName
}

type DepartmentName string

const (
    Finance   DepartmentName = "finance"
    Warehouse DepartmentName = "warehouse"
    Marketing DepartmentName = "marketing"
)
```

Struct `InvoiceData` akan digunakan untuk menyimpan data invoice yang formatnya sudah bisa digunakan untuk melakukan rekapitulasi data. Terdapat data berupa tanggal, total harga, dan departemetn mana.

- `Date` bertipe `string` berformat `"DD-month-YYYY"` (contoh: "01-january-2022")
- `TotalInvoice` bertipe `float64` yang akan berisi total harga dari invoice finance di data `Details`
- `DepartmentName` akan berisi informasi departemen mana yang menghasilkan invoice tersebut ("finance", "warehouse", dan "marketing").

Return ke 2 adalah `error`, lakukan handling error jika terdapat data yang tidak sesuai sebagai berikut:

- Jika `Date` berisi _string_ kosong `"DD-month-YYYY"` maka return `error` dengan pesan `"invoice date is format"`
- Jika isi dari `Details` kosong maka return `error` dengan pesan `"invoice details is empty"`
- Jika `Status` di invoice finance berisi string kosong atau selain "paid" dan "unpaid", maka return `error` dengan pesan `"invoice status is not valid"`
- Jika terdapat `Total` di `Details` yang bernilai 0 atau negatif, maka return `error` dengan pesan `"total price is not valid"`

## Phase 2

Kita akan mengimplementasikan _interface_ untuk _struct_ inovice departemen warehouse. Buatlah method yang dibutuhkan agar dapat melakukan konversi data invoice warehouse ke format yang sebelumnya sudah dijelaskan.

Untuk `TotalInvoice` dari invoice warehouse, diakumulasikan pada setiap product dengan rumus sebagai berikut:

```txt
total x price - discout price
```

Contoh, jika terdapat 2 product dengan total 2 dan 3, harga 10000 dan 20000, dan diskon `0.1` dan `0.2`, maka total invoice adalah:

- product pertama, `2 x 10000 = 18000` dikurangi diskon `0.1` atau 10% (`1800`) maka totalnya adalah `16200`
- product kedua, `3 x 20000 = 60000` dikurangi diskon `0.2` atau 20% (`12000`) maka totalnya adalah `48000`

Maka total invoice adalah `16200 + 48000 = 64200`

Jangan lupa untuk melakukan handling error jika terdapat data yang tidak sesuai sebagai berikut:

- Jika `Date` berisi _string_ kosong `"DD-month-YYYY"` maka return `error` dengan pesan `"invoice date is format"`
- Jika `InvoiceType` berisi string kosong atau selain "purchase" dan "sales", maka return `error` dengan pesan `"invoice type is invalid"`
- Jika isi dari `Products` kosong maka return `error` dengan pesan `"invoice products is empty"`
- Jika terdapat total pembelian product (`Unit`) yang bernilai 0 atau kurang dari 0, maka return `error` dengan pesan `"unit product is not valid"`
- Jika terdapat harga product (`Price`) yang bernilai 0 atau kurang dari 0, maka return `error` dengan pesan `"price product is not valid"`

### Phase 3

Kita akan mengimplementasikan _interface_ untuk _struct_ inovice departemen marketing. Buatlah method yang dibutuhkan agar dapat melakukan konversi data invoice marketing ke format yang sebelumnya sudah dijelaskan.

Untuk `TotalInvoice` dari invoice marketing, diakumulasikan dari lama perjalanan (dalam hari) dikali dengan biaya per hari dan ditambah dengan biaya lainnya.

```txt
(end date - start date) x price per day + another fee
```

Contoh, jika `StartDate` adalah "01-january-2022" dan `EndDate` adalah `"03-january-2022"`, maka lama perjalanan adalah 3 hari. Jika `PricePerDay` adalah 10000 dan `AnotherFee` adalah 5000, maka total invoice adalah 35000.

Jangan lupa untuk melakukan handling error jika terdapat data yang tidak sesuai sebagai berikut:

- Jika `Date` berisi _string_ kosong, maka return `error` dengan pesan `"invoice date is empty"`
- Jika `StartDate` atau `EndDate` berisi string kosong, maka return `error` dengan pesan `"travel date is empty"`
- Jika `PricePerDay` bernilai 0 atau kurang dari 0, maka return `error` dengan pesan `"price per day is not valid"`

### Phase 4

Kita akan gunakan _interface_ ke dalam fungsi `RecapDataInvoince` untuk menghitung total invoice dari setiap departemen. Buatlah fungsi `RecapDataInvoince` yang menerima parameter `invoices` bertipe `[]Invoice` (_interface_). FUngsi ini akan mengembalikan _array_ `InvoiceData` dan `error`. Kita dapat menggunakan method dari `RecordInvoice` untuk mengubahnya menjadi `InvoiceData`. Jangan lupa untuk melakukan handling error jika terdapat data yang tidak sesuai dari invoice ke 3 departement yang sudah dijelaskan sebelumnya.

```go
func RecapDataInvoice(data []Invoice) ([]InvoiceData, error) {
    // kerjakan di sini
}
```

FUngsi ini akan mengumpulkan semua invoice dengan ketentuan sebagai berikut:

- Hanya akan mengumpulkan invoice yang `Status` nya adalah "paid" (khusus untuk invoice finance)
- Hanya akan mengumpulkan `InvoiceType` nya adalah "purchsae" (khusus untuk invoice Warehouse)
- `Approved` nya bernilai `true` atau sudah ditandatangani
- Menggabungkan invoice dengan tanggal yang sama dan departemen yang sama menjadi satu data.

Contoh, jika terdapat 3 invoice sebagai berikut:

```txt
data 1 : invoice finance 
[
    {
        Date: "01/01/2022",
        Details: [
            {
                "Description": "pembelian nota",
                "Total": 40000
            },
            {
                "Description": "pembelian alat tulis",
                "Total": 20000
            }
        ]
        Approved: true
        Status: "paid"
    },
    {
        Date: "01/01/2022",
        Details: [
            {
                "Description": "pembelian nota 2",
                "Total": 50000
            },
            {
                "Description": "pembelian peralatan",
                "Total": 100000
            }
        ]
        Approved: true
        Status: "paid"
    },
    {
        Date: "01/01/2022",
        Details: [
            {
                "Description": "pembelian nota 3",
                "Total": 50000
            },
            {
                "Description": "pembelian peralatan",
                "Total": 100000
            }
        ]
        Approved: false
        Status: "unpaid"
    }
]
```

Maka Invoice yang akan dihasilkan adalah sebagai berikut:

```txt
[
    {
        Date: "01/01/2022",
        Department: "finance",
        Total: 210000
    }
]
```

Karena ketiga invoice tersebut memiliki tanggal yang sama dan departemen yang sama, maka akan dijumlahkan menjadi satu data. Dan yang dihitung hanya berstatus `"paid"` dan `Approved` yang sudah ditandatangani `true`.

### Test Case Examples

#### Test Case 1

**Input**:

Kita menggunakan fungsi `RecapDataInvoice` untuk melakukan pengujian ini

```txt
data = [
    // invoice Finance
    {
        Date: "01/02/2020", 
        Details: [
            {
                "Description": "pembelian nota",
                "Total": 4000
            },
            {
                "Description": "pembelian alat tulis",
                "Total": 4000
            }
        ], 
        Approved: true, Status: "paid"
    },
    {
        Date: "01/02/2020", 
        Details: [
            {
                "Description": "pembelian nota",
                "Total": 4000
            },
            {
                "Description": "pembelian alat tulis",
                "Total": 4000
            }
        ], 
        Approved: true, Status: "paid"
    },

    // invoice warehouse
    {
        Date: "01-February-2020",
        Products: [
            {
                Name: "Beras",
                Unit: 10,
                Price: 10000,
                Quantity: 0.1,
            },
            {
                Name: "Gula",
                Unit: 5,
                Price: 15000,
                Discout: 0.2
            },
            }

        ],
        InvoiceType: "purchase",
        Approved:    true,
    },

    // invoice marketing
    {
        Date: "01/02/2020",
        StartDate: 20/01/2020",
        EndDate: "25/01/2020",
        PricePerDay: 10000,
        AnotherFree: 5000,
        Approved: true,
    },
]
```

**Expected Output / Behavior**:

```txt
InvoiceData = [
    {
        Date: "01-February-2020",
        Department: "finance",
        Total: 16000
    },
    {
        Date: "01-February-2020",
        Department: "warehouse",
        Total: 150000
    },
    {
        Date: "01-February-2020",
        Department: "marketing",
        Total: 65000
    }
]

error = nil
```

**Explanation**:

Seperti yang dijelaskan sebelumnya bahwa fungsi ini akan mengumpulkan semua invoice dengan tanggal yang sama dan departemen yang sama menjadi satu data. Dan yang dihitung hanya berstatus `"paid"` dan `Approved` yang sudah ditandatangani `true`. Selama tidak terjadi error maka akan mengembalikan `InvoiceData` seperti contoh di atas

#### Test Case 2

**Input**:

Kita menggunakan fungsi `RecapDataInvoice` untuk melakukan pengujian ini

```txt
data = [{
    // invoice Finance
    {
        Date: "", 
        Details: [
            {
                "Description": "pembelian nota",
                "Total": 4000
            },
            {
                "Description": "pembelian alat tulis",
                "Total": 4000
            }
        ], 
        Approved: true, Status: "paid"
    },
}]
```

**Expected Output / Behavior**:

```txt
InvoiceData = nil

error = "invoice date is empty"
```

**Explanation**:

Karena invoice date kosong maka akan mengembalikan _error_ `"invoice date is empty"`.
