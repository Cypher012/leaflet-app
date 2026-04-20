import { Button } from '#/components/ui/button'
import useOAuth from '#/features/auth/hooks/auth.useOAuth'
import { GithubIcon } from '#/icons/github.icon'
import { GoogleIcon } from '#/icons/google.icon'

export const OAuthButtons = () => {
  const { loginWithGitHub, loginWithGoogle } = useOAuth()
  return (
    <div className="flex flex-col gap-4 w-full">
      <Button
        variant={'secondary'}
        className="py-6 rounded-4xl"
        type="button"
        onClick={loginWithGoogle}
      >
        <GoogleIcon className="mx-2" />
        Continue with Google
      </Button>
      <Button
        variant={'secondary'}
        className="py-6 rounded-4xl"
        type="button"
        onClick={loginWithGitHub}
      >
        <GithubIcon className="mx-2" />
        Continue with GitHub
      </Button>
    </div>
  )
}
