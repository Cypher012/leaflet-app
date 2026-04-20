import CreatePage from '#/features/create/components/createPage'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/__protected/create/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <CreatePage />
}
