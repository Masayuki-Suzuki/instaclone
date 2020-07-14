import React from 'react'
import { css } from '@emotion/core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { IconName } from '@fortawesome/fontawesome-common-types'

type SnsLoginButtonPropsType = {
    brand?: IconName
    validation: boolean
    text: string
    onClickButton?: () => void
}

const wrapper = css({
    width: '100%'
})

const button = css({
    background: '#0095f6',
    border: '1px solid transparent',
    borderRadius: '4px',
    color: '#fff',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    fontSize: '1.4rem',
    fontWeight: 500,
    letterSpacing: '.02em',
    lineHeight: '18px',
    padding: '6px 9px',
    width: '100%'
})

const icon = css({
    marginRight: '8px'
})

const SnsLoginButton = ({ brand, onClickButton, text, validation }: SnsLoginButtonPropsType): JSX.Element => {
    return (
        <div css={wrapper}>
            <button
                onClick={onClickButton && (() => onClickButton())}
                css={css`
                    ${button};
                    opacity: ${validation ? 1 : 0.4};
                `}
            >
                {brand ? <FontAwesomeIcon css={icon} icon={['fab', brand]} /> : null}
                {text}
            </button>
        </div>
    )
}

export default SnsLoginButton
