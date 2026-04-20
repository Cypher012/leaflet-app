import FeedPage from '#/features/profile/components/page/feedPage'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/profile/$username/posts')({
  component: RouteComponent,
})

function RouteComponent() {
  const { username } = Route.useParams()
  return <FeedPage username={username} />
}
