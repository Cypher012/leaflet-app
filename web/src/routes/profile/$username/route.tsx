import ProfilePage from '#/features/profile/components/profilePage'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/profile/$username')({
  component: RouteComponent,
})

function RouteComponent() {
  const { username } = Route.useParams()

  return (
    <ProfilePage username={username}>
      <Outlet />
    </ProfilePage>
  )
}
