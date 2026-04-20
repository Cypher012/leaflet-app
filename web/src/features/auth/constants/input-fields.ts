type LoginFieldConfig = {
  name: 'email' | 'password'
  label: string
  placeholder?: string
  type?: string
  forgetPassword?: boolean
}

type SignupFieldConfig = {
  name: 'fullName' | 'email' | 'password' | 'confirmPassword'
  label: string
  placeholder?: string
  type?: string
}[]

export const loginFields: LoginFieldConfig[] = [
  {
    name: 'email',
    label: 'EMAIL ADDRESS',
    placeholder: 'curator@leaflet.com',
  },
  {
    name: 'password',
    label: 'PASSWORD',
    type: 'password',
    placeholder: '********',
    forgetPassword: true,
  },
]

export const signupFields: SignupFieldConfig[] = [
  [
    {
      name: 'fullName',
      label: 'FULL NAME',
      placeholder: 'John Doe',
    },
  ],
  [
    {
      name: 'email',
      label: 'EMAIL ADDRESS',
      placeholder: 'curator@leaflet.com',
    },
  ],
  [
    {
      name: 'password',
      label: 'PASSWORD',
      type: 'password',
      placeholder: '********',
    },
    {
      name: 'confirmPassword',
      label: 'CONFIRM PASSWORD',
      type: 'password',
      placeholder: '********',
    },
  ],
]
