import { Link, useRouteContext } from '@tanstack/react-router'
import { buttonVariants } from '../ui/button'
import { cn } from '#/lib/utils'
import { Logo } from '#/icons/logo.icon'
import UserMenu from './UserMenu'
import { useEffect, useState } from 'react'
import LoginModal from '#/features/auth/components/login-modal'
import SignupModal from '#/features/auth/components/signup-modal'

const Header = () => {
  const { user } = useRouteContext({ from: '__root__' })
  const [loginOpen, setLoginOpen] = useState(false)
  const [signupOpen, setSignupOpen] = useState(false)

  useEffect(() => {
    const handler = () => setLoginOpen(true)
    window.addEventListener('auth:unauthorized', handler)
    return () => window.removeEventListener('auth:unauthorized', handler)
  }, [])

  return (
    <header
      suppressHydrationWarning
      className="flex justify-between items-center py-4 px-6 md:px-8 sticky top-0 z-50 border-b backdrop-blur border-border/20 supports-backdrop-filter:bg-background/60"
    >
      <Link to="/">
        <Logo className="scale-90" />
      </Link>

      <div className="space-x-5 md:flex hidden items-center">
        {user ? (
          <UserMenu user={user} />
        ) : (
          <button
            onClick={() => setLoginOpen(true)}
            className={cn(
              buttonVariants({ variant: 'default' }),
              'py-5 px-5 rounded-3xl',
            )}
          >
            Log in
          </button>
        )}
      </div>

      <LoginModal
        open={loginOpen}
        onOpenChange={setLoginOpen}
        onSignup={() => setSignupOpen(true)}
      />
      <SignupModal open={signupOpen} onOpenChange={setSignupOpen} />
    </header>
  )
}

export default Header
