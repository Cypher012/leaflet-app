import { Link } from '@tanstack/react-router'

export default function NotFound() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen gap-4">
      <h1 className="text-4xl font-bold">404</h1>
      <p className="text-muted-foreground">This page doesn't exist.</p>
      <Link to="/">Go home</Link>
    </div>
  )
}
