import CommentPage from '#/features/profile/components/page/commentPage'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/profile/$username/comments')({
  component: RouteComponent,
})

function RouteComponent() {
  const { username } = Route.useParams()
  return <CommentPage username={username} />
}
