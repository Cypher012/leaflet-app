import {
  Field,
  FieldError,
  FieldGroup,
  FieldLabel,
} from '#/components/ui/field'
import { Input } from '#/components/ui/input'
import { Link } from '@tanstack/react-router'

type FormInputProps = {
  form: any
  name: string
  label: string
  placeholder?: string
  type?: string
  forgetPassword?: boolean
}

function FormInput({
  form,
  name,
  label,
  forgetPassword = false,
  placeholder,
  type = 'text',
}: FormInputProps) {
  return (
    <FieldGroup>
      <form.Field
        name={name}
        children={(field: any) => {
          const isInvalid =
            field.state.meta.isTouched && !field.state.meta.isValid

          return (
            <Field data-invalid={isInvalid} className="relative">
              <FieldLabel
                className="text-gray-500 text-xs"
                htmlFor={field.name}
              >
                {label}
                {forgetPassword && (
                  <Link
                    to="/"
                    className="text-xs text-gray-500 absolute right-0 top-0"
                  >
                    FORGET PASSWORD?
                  </Link>
                )}
              </FieldLabel>

              <Input
                id={field.name}
                name={field.name}
                type={type}
                value={field.state.value}
                onBlur={field.handleBlur}
                onChange={(e) => field.handleChange(e.target.value)}
                aria-invalid={isInvalid}
                placeholder={placeholder}
                autoComplete="off"
                className="shadow-none! border-none bg-muted py-7! px-4!"
              />

              {isInvalid && <FieldError errors={field.state.meta.errors} />}
            </Field>
          )
        }}
      />
    </FieldGroup>
  )
}

export default FormInput
