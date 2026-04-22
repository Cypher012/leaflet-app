// src/components/web/UserMenu.tsx
import { useState } from 'react'
import { motion, AnimatePresence, easeInOut } from 'framer-motion'
import { Home, LogOut, Settings, SquarePen, User } from 'lucide-react'
import { Link } from '@tanstack/react-router'
import ThemeToggle from './ThemeToggle'
import { useLogout } from '#/features/auth/hooks/uselogout'

interface UserMenuProps {
  user: {
    fullname: string
    username: string
    avatar_url: string
  }
}

const menuVariants = {
  closed: {
    opacity: 0,
    x: 20,
    scale: 0.95,
    transition: { duration: 0.2, ease: easeInOut },
  },
  open: {
    opacity: 1,
    x: 0,
    scale: 1,
    transition: { duration: 0.3, ease: easeInOut, staggerChildren: 0.05 },
  },
}

const itemVariants = {
  closed: { opacity: 0, x: 10 },
  open: { opacity: 1, x: 0 },
}

export default function UserMenu({ user }: UserMenuProps) {
  const [isOpen, setIsOpen] = useState(false)
  const { logout } = useLogout()
  return (
    <div className="relative">
      <motion.button
        onClick={() => setIsOpen(!isOpen)}
        className="flex items-center focus:outline-none"
        whileTap={{ scale: 0.95 }}
      >
        <div className="md:size-9 size-8  overflow-hidden rounded-full border-2 border-transparent transition-colors hover:border-primary">
          <img
            src={user.avatar_url}
            alt={user.fullname}
            className="h-full w-full object-cover"
          />
        </div>
      </motion.button>

      <AnimatePresence>
        {isOpen && (
          <>
            <motion.div
              className="fixed inset-0 z-40 bg-transparent"
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              exit={{ opacity: 0 }}
              onClick={() => setIsOpen(false)}
            />

            <motion.div
              className="border-border/30 bg-background absolute right-0 z-50 mt-3 md:w-64 w-55 overflow-hidden rounded-2xl border shadow-2xl"
              variants={menuVariants}
              initial="closed"
              animate="open"
              exit="closed"
            >
              <div className="p-4">
                <div className="mb-4 flex flex-col px-2">
                  <span className="text-foreground truncate md:text-sm text-xs font-bold">
                    {user.fullname}
                  </span>
                  <span className="text-muted-foreground truncate md:text-xs text-[10px]">
                    {user.username}
                  </span>
                </div>

                <div className="space-y-1">
                  <UserMenuItem
                    href="/"
                    icon={<Home className="md:size-4 size-3.5" />}
                    label="Home"
                    onClick={() => setIsOpen(false)}
                  />
                  <UserMenuItem
                    href="/profile/$username"
                    params={{username: user.username}}
                    icon={<User className="md:size-4 size-3.5" />}
                    label="Profile"
                    onClick={() => setIsOpen(false)}
                  />
                  <UserMenuItem
                    href="/create"
                    icon={<SquarePen className="md:size-4 size-3.5" />}
                    label="Create"
                    onClick={() => setIsOpen(false)}
                  />
                  <UserMenuItem
                    href="/settings"
                    icon={<Settings className="md:size-4 size-3.5" />}
                    label="Settings"
                    onClick={() => setIsOpen(false)}
                  />
                  <motion.div variants={itemVariants}>
                    <ThemeToggle />
                  </motion.div>
                </div>

                <motion.div
                  className="border-border mt-4 border-t pt-2"
                  variants={itemVariants}
                >
                  <button
                    onClick={logout}
                    className="md:text-sm text-xs text-destructive hover:bg-destructive/10 flex w-full items-center space-x-3 rounded-lg px-3 py-2 font-medium transition-colors"
                  >
                    <LogOut className="md:size-4 size-3.5" />
                    <span>Sign Out</span>
                  </button>
                </motion.div>
              </div>
            </motion.div>
          </>
        )}
      </AnimatePresence>
    </div>
  )
}

function UserMenuItem({
  href,
  icon,
  label,
  params,
  onClick,
}: {
  href: string
  icon: React.ReactNode
  params?: Record<string, string>
  label: string
  onClick: () => void
}) {
  return (
    <motion.div variants={itemVariants}>
      <Link
        to={href}
        onClick={onClick}
        params={params}
        className="md:text-sm text-xs text-foreground/80 hover:bg-muted hover:text-foreground flex items-center space-x-3 rounded-lg px-3 py-2 font-medium transition-colors"
      >
        {icon}
        <span>{label}</span>
      </Link>
    </motion.div>
  )
}
