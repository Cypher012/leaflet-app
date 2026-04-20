import { loginSchema } from '#/features/auth/schemas/auth'
import { useForm } from '@tanstack/react-form'

const useSignupForm = () => {
  const form = useForm({
    defaultValues: {
      email: '',
      password: '',
    },
    validators: {
      onSubmit: loginSchema,
    },
    onSubmit: (values) => {
      console.log(values)
    },
  })

  return form
}

export default useSignupForm
