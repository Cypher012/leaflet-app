import { signUpSchema } from '#/features/auth/schemas/auth'
import { useForm } from '@tanstack/react-form'

const useSignupForm = () => {
  const form = useForm({
    defaultValues: {
      fullName: '',
      email: '',
      password: '',
      confirmPassword: '',
    },
    validators: {
      onSubmit: signUpSchema,
    },
    onSubmit: (values) => {
      console.log(values)
    },
  })

  return form
}

export default useSignupForm
