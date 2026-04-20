import { Button } from '#/components/ui/button'
import FormInput from '#/components/web/FormInput'
import useLoginForm from '#/features/auth/hooks/login.useform'
import { Link } from '@tanstack/react-router'
import { ArrowRight } from 'lucide-react'
import { loginFields as fields } from '../constants/input-fields'
import { Logo } from '#/icons/logo.icon'
import { OAuthButtons } from './oauth-buttons'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from '#/components/ui/dialog'

type LoginModalProps = {
  open: boolean
  onOpenChange: (open: boolean) => void
  onSignup?: () => void
}

const LoginModal = ({ open, onOpenChange, onSignup }: LoginModalProps) => {
  const form = useLoginForm()

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-xl">
        <DialogHeader className="flex flex-col items-center">
          <Link to="/">
            <Logo className="scale-125 mx-auto mb-6 mt-2" />
          </Link>
          <DialogTitle className="text-2xl font-medium">
            Welcome back
          </DialogTitle>
          <DialogDescription asChild>
            <p className="text-foreground">
              Don't have an account?{' '}
              <button
                onClick={() => {
                  onOpenChange(false)
                  onSignup?.()
                }}
                className="text-primary font-medium"
              >
                Sign up
              </button>
            </p>
          </DialogDescription>
        </DialogHeader>

        <div className="flex flex-col space-y-10 mt-2">
          <div className="w-full flex flex-col items-center">
            <OAuthButtons />
            <div className="flex items-center gap-4 mt-8 w-full">
              <div className="flex-1 border-t border-border" />
              <span className="text-xs text-muted-foreground uppercase tracking-wide">
                Or continue with email
              </span>
              <div className="flex-1 border-t border-border" />
            </div>
          </div>

          <form
            onSubmit={(e) => {
              e.preventDefault()
              form.handleSubmit()
            }}
          >
            <div className="flex flex-col space-y-6">
              {fields.map((field) => (
                <FormInput key={field.name} form={form} {...field} />
              ))}
              <Button className="py-6 mt-3" type="submit">
                Log In <ArrowRight />
              </Button>
            </div>
          </form>
        </div>
      </DialogContent>
    </Dialog>
  )
}

export default LoginModal
