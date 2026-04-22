import { Link } from '@tanstack/react-router'
import { Avatar, AvatarFallback, AvatarImage } from '#/components/ui/avatar'
import type React from 'react'
import type { ProfileUserResponse } from '../types'

const tabs = [
  { label: 'Overview', to: '/', exact: true },
  { label: 'Posts', to: '/posts', exact: false },
  { label: 'Comments', to: '/comments', exact: false },
]

const MainContent = ({
  profile,
  children,
}: {
  profile: ProfileUserResponse
  children: React.ReactNode
}) => {
  return (
    <div className="p-3">
      <div className="flex gap-x-5 items-center">
        <Avatar className="md:size-20 size-14">
          <AvatarImage src={profile.avatar_url} alt="avatar_url" />
          <AvatarFallback>{profile.fullname}</AvatarFallback>
        </Avatar>
        <div className="space-y-1">
          <p className="font-semibold text-foreground md:text-3xl text-xl">
            {profile.fullname}
          </p>
          <p className="text-muted-foreground md:text-base text-sm">{profile.username}</p>
        </div>
      </div>
      <div className="flex gap-x-3 mt-12 font-medium">
        {tabs.map((tab) => (
          <Link
            key={tab.to}
            to={'/profile/$username' + tab.to}
            params={{username: profile.username}}
            activeOptions={{ exact: tab.exact }}
            activeProps={{
              className:
                'bg-primary text-primary-foreground rounded-md px-4 py-2 md:text-sm text-xs',
            }}
            inactiveProps={{
              className:
                'bg-secondary text-secondary-foreground rounded-md px-4 py-2 md:text-sm text-xs',
            }}
          >
            {tab.label}
          </Link>
        ))}
      </div>
      <div className="mt-6">{children}</div>
    </div>
  )
}

export default MainContent
