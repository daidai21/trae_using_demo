# E-Commerce Platform

A complete e-commerce trading platform developed from scratch, adopting a frontend-backend separated architecture.

## Tech Stack

### Backend
- **Language**: Golang
- **Web Framework**: Hertz
- **ORM**: GORM
- **Database**: SQLite
- **Authentication**: JWT
- **Password Encryption**: bcrypt

### Frontend
- **Framework**: React 18
- **UI Component Library**: Ant Design 5
- **Routing**: React Router v6
- **HTTP Client**: Axios
- **Build Tool**: Vite
- **Language**: TypeScript

## Project Structure

```
trae_using_demo/
├── backend/              # Backend service
│   ├── cmd/             # Application entry point
│   │   └── server/
│   ├── internal/        # Internal code
│   │   ├── handler/     # HTTP handlers
│   │   ├── service/     # Business logic layer
│   │   ├── repository/  # Data access layer
│   │   ├── model/       # Data models
│   │   └── middleware/  # Middleware
│   ├── pkg/             # Public packages
│   │   ├── response/    # Unified response format
│   │   └── utils/       # Utility functions
│   ├── bin/             # Compiled output
│   └── go.mod
├── frontend/            # Frontend application
│   ├── src/
│   │   ├── components/  # Public components
│   │   ├── pages/       # Page components
│   │   ├── services/    # API services
│   │   ├── store/       # State management
│   │   ├── types/       # TypeScript types
│   │   ├── App.tsx
│   │   └── main.tsx
│   └── package.json
└── README.md
```

## Features

### User Authentication
- User registration
- User login
- JWT Token authentication (24-hour validity)
- Password bcrypt encryption storage

### Merchant Management
- Merchant registration
- Merchant information editing
- Merchant list query
- Merchant detail view

### Product Management
- Product publishing
- Product list (public access)
- Product details
- Product editing
- Product deletion
- Inventory management

### Shopping Cart
- Add products to cart
- Cart list
- Modify product quantity
- Delete cart items

### Order Management
- Create orders (transaction processing)
- Order list
- Order details
- Order status update
- Automatic inventory deduction
- Automatic cart clearing

## Quick Start

### Backend Startup

```bash
cd backend
go build -o bin/server ./cmd/server
./bin/server
```

The backend service will start at `http://localhost:8080`

### Frontend Startup

```bash
cd frontend
npm install
npm run dev
```

The frontend development server will start at `http://localhost:5173`

## API Documentation

### Authentication Endpoints
- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login

### Product Endpoints (Public)
- `GET /api/products` - Get product list
- `GET /api/products/:id` - Get product details

### Authenticated Endpoints

#### Merchant Management
- `POST /api/merchants` - Create merchant
- `GET /api/merchants` - Get merchant list
- `GET /api/merchants/:id` - Get merchant details
- `PUT /api/merchants/:id` - Update merchant information

#### Product Management
- `POST /api/products` - Create product
- `PUT /api/products/:id` - Update product
- `DELETE /api/products/:id` - Delete product

#### Shopping Cart
- `POST /api/cart` - Add product to cart
- `GET /api/cart` - Get cart
- `PUT /api/cart/:id` - Update cart item
- `DELETE /api/cart/:id` - Delete cart item

#### Order Management
- `POST /api/orders` - Create order
- `GET /api/orders` - Get order list
- `GET /api/orders/:id` - Get order details
- `PUT /api/orders/:id/status` - Update order status

## Database Schema

### users (User Table)
- id: Primary key
- username: Username (unique)
- password: Password (encrypted)
- created_at: Creation time
- updated_at: Update time
- deleted_at: Deletion time (soft delete)

### merchants (Merchant Table)
- id: Primary key
- user_id: User ID (foreign key)
- name: Merchant name
- description: Merchant description
- created_at: Creation time
- updated_at: Update time
- deleted_at: Deletion time (soft delete)

### products (Product Table)
- id: Primary key
- merchant_id: Merchant ID (foreign key)
- name: Product name
- description: Product description
- price: Product price
- stock: Inventory quantity
- created_at: Creation time
- updated_at: Update time
- deleted_at: Deletion time (soft delete)

### carts (Shopping Cart Table)
- id: Primary key
- user_id: User ID (foreign key)
- product_id: Product ID (foreign key)
- quantity: Quantity
- created_at: Creation time
- updated_at: Update time
- deleted_at: Deletion time (soft delete)

### orders (Order Table)
- id: Primary key
- user_id: User ID (foreign key)
- total_amount: Order total amount
- status: Order status
- created_at: Creation time
- updated_at: Update time
- deleted_at: Deletion time (soft delete)

### order_items (Order Item Table)
- id: Primary key
- order_id: Order ID (foreign key)
- product_id: Product ID (foreign key)
- quantity: Quantity
- price: Product unit price
- created_at: Creation time
- updated_at: Update time
- deleted_at: Deletion time (soft delete)

## Development Notes

- Backend uses Hertz framework with GORM for data persistence
- Frontend uses React + Ant Design with React Router for routing
- Uses JWT for identity authentication with 24-hour token validity
- Passwords are encrypted and stored using bcrypt
- Transaction processing is used when creating orders to ensure data consistency
- Supports CORS cross-origin requests

## License

MIT License
