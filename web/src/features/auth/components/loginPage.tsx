import { Button } from '#/components/ui/button'
import FormInput from '#/components/web/FormInput'
import useLoginForm from '#/features/auth/hooks/login.useform'
import { Link } from '@tanstack/react-router'
import { ArrowRight } from 'lucide-react'
import { loginFields as fields } from '../constants/input-fields'
import { Logo } from '#/icons/logo.icon'
import BackButton from './back-button'
import { OAuthButtons } from './oauth-buttons'

const LoginPage = () => {
  const form = useLoginForm()

  return (
    <div className="relative">
      <BackButton />
      <div className="section-container mx-auto max-w-xl pt-16">
        <Link to="/">
          <Logo className="scale-125 mx-auto mb-10 mt-5" />
        </Link>

        <div className="flex flex-col space-y-10">
          <div className="space-y-2.5 flex justify-center flex-col items-center">
            <h1 className="font-medium text-3xl">Welcome back</h1>
            <p className="text-foreground">
              Don't have an account?{' '}
              <Link className="text-primary font-medium" to="/signup">
                Sign up
              </Link>
            </p>
          </div>

          {/* OAUTH FIRST */}
          <div className="w-full flex flex-col items-center">
            <OAuthButtons />

            {/* Divider */}
            <div className="flex items-center gap-4 mt-8 w-full">
              <div className="flex-1 border-t border-border" />

              <span className="text-xs text-muted-foreground uppercase tracking-wide">
                Or continue with email
              </span>

              <div className="flex-1 border-t border-border" />
            </div>
          </div>

          {/* MANUAL FORM AFTER */}
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
      </div>
    </div>
  )
}

export default LoginPage
