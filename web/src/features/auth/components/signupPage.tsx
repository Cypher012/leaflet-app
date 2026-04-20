import { Button } from '#/components/ui/button'
import FormInput from '#/components/web/FormInput'
import useSignupForm from '#/features/auth/hooks/signup.useform'
import { Link } from '@tanstack/react-router'
import { ArrowRight } from 'lucide-react'
import { signupFields as fields } from '../constants/input-fields'
import { Logo } from '#/icons/logo.icon'
import { GoogleIcon } from '#/icons/google.icon'
import { GithubIcon } from '#/icons/github.icon'
import BackButton from './back-button'
import { OAuthButtons } from './oauth-buttons'

const SignupPage = () => {
  const form = useSignupForm()

  return (
    <div className="relative">
      <BackButton />
      <div className="section-container mx-auto max-w-xl pt-20 lg:pt-16 ">
        <Link to="/">
          <Logo className="scale-125 mx-auto mb-10 mt-5" />
        </Link>

        <div className="flex flex-col space-y-10">
          <div className="space-y-2.5">
            <h1 className="font-medium text-3xl">Create your account</h1>
            <p className="text-foreground">
              Already have an account?{' '}
              <Link className="text-primary font-medium" to="/login">
                Log in
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
      </div>
    </div>
  )
}

export default SignupPage
