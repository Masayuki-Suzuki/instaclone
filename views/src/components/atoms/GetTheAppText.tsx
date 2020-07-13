import React from 'react'
import { css } from '@emotion/core'

const text = css({
    color: '#262626',
    fontSize: '1.4rem',
    lineHeight: '18px',
    textAlign: 'center',
    margin: '10px 20px'
})

const GetTheAppText = (): JSX.Element => {
    return <p css={text}>Get the app.</p>
}

export default GetTheAppText
