# Go-HTMX with Tailwind CSS v4

A simple Go web application using HTMX and Tailwind CSS v4.

## Setup

1. Install Node.js dependencies:
```bash
npm install
```

2. Build Tailwind CSS:
```bash
npm run build:css
```

3. Run the Go server:
```bash
go run main.go
```

4. Open http://localhost:8000 in your browser

## Development

To watch for CSS changes during development:
```bash
npm run watch:css
```

## Scripts

- `npm run build:css` - Build Tailwind CSS once
- `npm run watch:css` - Watch for changes and rebuild automatically
