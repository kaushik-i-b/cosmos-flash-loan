package types

type FlashLoanRequest struct {
    Lender   string  `json:"lender"`
    Borrower string  `json:"borrower"`
    Amount   int64   `json:"amount"`
    Fee      float64 `json:"fee"`
}

type FlashLoanResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
}

type LoanState struct {
    IsActive bool   `json:"is_active"`
    Amount   int64  `json:"amount"`
    DueDate  string `json:"due_date"`
}