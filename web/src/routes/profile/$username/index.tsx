import OverviewPage from '#/features/profile/components/page/overviewPage'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/profile/$username/')({
  component: RouteComponent,
})

function RouteComponent() {
  const { username } = Route.useParams()
  return <OverviewPage username={username} />
}