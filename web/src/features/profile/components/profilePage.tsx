import MainContent from './mainContent'
import SideContent from './sideContent'
import { useUserProfileQuery } from '../api/queries'
import ProfilePageSkeleton from './skeleton/profile-page-skeleton'

const ProfilePage = ({ children, username }: { children: React.ReactNode, username: string }) => {
  const {
    data: profile
  } = useUserProfileQuery(username)

  if (!profile) {
    return <ProfilePageSkeleton/>
  }

  return (
    <div className="flex w-full gap-3 max-w-7xl mx-auto py-10 px-3 sm:px-6 lg:px-4">
      <div className="w-full max-w-4xl lg:max-w-3xl mx-auto">
        <MainContent profile={profile}>{children}</MainContent>
      </div>
      <div className="w-full max-w-xs mx-auto shrink-0 hidden lg:block">
        <div className="sticky bg-secondary/30 rounded-sm top-25 h-[30rem]">
          <SideContent profile={profile} />
        </div>
      </div>
    </div>
  )
}
export default ProfilePage
