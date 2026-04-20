import { createFileRoute, Outlet, redirect } from '@tanstack/react-router'

export const Route = createFileRoute('/__protected')({
  beforeLoad: ({ context, location }) => {
    if (!context.user) {
      throw redirect({
        to: '/',
        search: { redirect: location.href },
      })
    }
    return { user: context.user }
  },
  component: RouteComponent,
})

function RouteComponent() {
  return <Outlet />
}
