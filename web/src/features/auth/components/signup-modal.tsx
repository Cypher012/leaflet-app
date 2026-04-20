import { Button } from '#/components/ui/button'
import FormInput from '#/components/web/FormInput'
import useSignupForm from '#/features/auth/hooks/signup.useform'
import { Link } from '@tanstack/react-router'
import { ArrowRight } from 'lucide-react'
import { signupFields as fields } from '../constants/input-fields'
import { Logo } from '#/icons/logo.icon'
import { OAuthButtons } from './oauth-buttons'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
} from '#/components/ui/dialog'

type SignupModalProps = {
  open: boolean
  onOpenChange: (open: boolean) => void
}

const SignupModal = ({ open, onOpenChange }: SignupModalProps) => {
  const form = useSignupForm()

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-xl">
        <DialogHeader className="flex flex-col items-center">
          <Link to="/">
            <Logo className="scale-125 mx-auto mb-6 mt-2" />
          </Link>
          <DialogTitle className="text-3xl font-medium">
            Create your account
          </DialogTitle>
          <DialogDescription asChild>
            <p className="text-foreground">
              Already have an account?{' '}
              <Link className="text-primary font-medium" to="/login">
                Log in
              </Link>
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
              {fields.map((row, i) => (
                <div
                  key={i}
                  className={`grid gap-4 ${row.length === 2 ? 'grid-cols-2' : 'grid-cols-1'}`}
                >
                  {row.map((field) => (
                    <FormInput key={field.name} form={form} {...field} />
                  ))}
                </div>
              ))}
              <Button className="py-6 mt-5" type="submit">
                Create account <ArrowRight />
              </Button>
            </div>
          </form>
        </div>
      </DialogContent>
    </Dialog>
  )
}

export default SignupModal
