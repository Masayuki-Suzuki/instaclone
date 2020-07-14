import React from 'react'
import { css } from '@emotion/core'
import ReduxFormInput from '~/components/atoms/ReduxFormInput'
import ValidationIcon from '~/components/atoms/ValidationIcon'

type ReduxFormInputIFieldPropsType = {
    input: unknown
    type: string
    placeholder: string
    meta: {
        touched: unknown
        error: unknown
        warning: unknown
    }
}

const wrapper = css({
    position: 'relative'
})

const ReduxFormInputIField = ({
    input,
    placeholder,
    type,
    meta: { touched, error }
}: ReduxFormInputIFieldPropsType): JSX.Element => {
    return (
        <div css={wrapper}>
            <ReduxFormInput input={input} placeholder={placeholder} type={type} />
            {touched && <ValidationIcon error={!!error} />}
        </div>
    )
}

export default ReduxFormInputIField
