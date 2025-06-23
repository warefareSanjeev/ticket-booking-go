
# Ticket-booking-GO

# 🎟️ GoLang Ticket Booking Application

This is a **basic ticket booking application** built with **Go (Golang)**. The project simulates a ticket booking system for a conference event, allowing users to register with their details and book available tickets concurrently using goroutines and wait groups.

## 🚀 Features

- Book tickets for a fictional conference
- Validate user input (name, email, ticket count)
- Track bookings and remaining tickets
- Send booking confirmation (simulated with goroutines)
- Modular code using helper packages
- Demonstrates Go fundamentals: structs, slices, loops, concurrency, and user input handling

## 🛠️ Technologies Used

- Go 1.20+
- Standard Go packages (`fmt`, `strings`, `sync`, `time`)

## 📂 Project Structure

booking_app/

├── main.go # Entry point of the application

├── helper/

│ └── input.go # Contains input validation logic


## ⚙️ How to Run

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/booking_app.git
   cd booking_app

2. Run the application:
go run main.go

✅ Validations
Name must be at least 2 characters

Email must contain @

Ticket number must be greater than 0 and less than remaining tickets

📌 Future Improvements
Store booking data in files or a database

Add a web interface using Go web frameworks (e.g. Gin or Echo)

Include automated tests

Add support for exporting ticket confirmations


