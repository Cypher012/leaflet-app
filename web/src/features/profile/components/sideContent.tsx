import type { ProfileUserResponse } from '../types'

const SideContent = ({ profile }: { profile: ProfileUserResponse }) => {
  return (
    <div className="p-6 space-y-5">
      <h2 className="text-xl font-semibold">{profile.fullname}</h2>
      <p className="text-sm text-muted-foreground">
        {profile.bio}
      </p>
      <div className="grid grid-cols-2 pt-6 gap-5">
        <p className="flex flex-col">
          <span className="text-secondary-foreground font-bold">{profile.stats.feed_count}</span>
          <span className="text-sm font-medium">Cards</span>
        </p>
        <p className="flex flex-col">
          <span className="text-secondary-foreground font-bold">{profile.stats.like_count}</span>
          <span className="text-sm font-medium">Likes</span>
        </p>
        <p className="flex flex-col">
          <span className="text-secondary-foreground font-bold">{profile.stats.comment_count}</span>
          <span className="text-sm font-medium">Comments</span>
        </p>
      </div>
    </div>
  )
}

export default SideContent
