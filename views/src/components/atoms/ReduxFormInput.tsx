import React from 'react'
import { css } from '@emotion/core'

type ReduxFormInputProps = {
    input: unknown
    placeholder: string
    type: string
}

const inputField = css({
    background: '#fafafa',
    border: 'solid 1px #dbdbdb',
    borderRadius: '3px',
    display: 'block',
    fontSize: '1.6rem',
    '-webkit-appearance': 'none',
    padding: '10px 0 10px 10px',
    width: '100%',
    '::placeholder': {
        color: '#aaa',
        fontSize: '12px',
        letterSpacing: '.04em'
    }
})

const ReduxFormInput = ({ input, placeholder, type }: ReduxFormInputProps): JSX.Element => {
    return <input {...input} css={inputField} type={type} placeholder={placeholder} />
}

export default ReduxFormInput
