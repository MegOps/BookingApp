# Conference Booking App

Welcome to the Conference Booking App! This command-line application is designed to manage bookings for a fictional conference by handling user information and ticket allocation.

## Features

- **Interactive Booking System**: Users can enter their personal details and request a number of tickets for the conference.
- **Concurrency**: Utilizes goroutines to send ticket confirmations in the background while processing further bookings.
- **Input Validation**: Ensures that users provide valid names, emails, and ticket numbers.
- **Real-time Updates**: Displays the number of remaining tickets and confirms each successful booking on the console.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.14 or higher recommended)

### Installation

1. Clone the repository to your local machine:
   ```bash
   git clone <repository-url>

### Based on:
https://www.youtube.com/watch?v=yyUHQIec83I
