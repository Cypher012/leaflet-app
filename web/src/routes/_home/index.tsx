import FeedsPage from '#/features/feeds/components/feedsPage'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_home/')({
  component: App,
})

function App() {
  return <FeedsPage />
}
