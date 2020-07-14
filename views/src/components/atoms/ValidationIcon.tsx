import React from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { css } from '@emotion/core'

type ValidationIconPropsType = {
    error: boolean
}

const base = css({
    fontSize: '2.2rem',
    fontWeight: 300,
    display: 'block',
    position: 'absolute',
    right: 10,
    top: '50%',
    transform: 'translateY(-50%)'
})

const ValidationIcon = ({ error }: ValidationIconPropsType): JSX.Element => {
    return (
        <React.Fragment>
            {error ? (
                <FontAwesomeIcon css={base} icon={['far', 'times-circle']} color="#d22" />
            ) : (
                <FontAwesomeIcon css={base} icon={['far', 'check-circle']} color="#aaa" />
            )}
        </React.Fragment>
    )
}

export default ValidationIcon
