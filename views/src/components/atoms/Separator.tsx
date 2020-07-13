import React from 'react'
import { css } from '@emotion/core'

const wrapper = css({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    padding: '10px 40px 18px',
    width: '100%'
})

const line = css({
    flexGrow: 1,
    flexShrink: 1,
    background: '#ccc',
    height: '1px'
})

const text = css({
    background: '#fff',
    color: '#8e8e8e',
    fontSize: '1.3rem',
    fontWeight: 500,
    letterSpacing: '.04em',
    lineHeight: '15px',
    textTransform: 'uppercase',
    flexGrow: 0,
    flexShrink: 0,
    margin: '0 18px'
})

const Separator = (): JSX.Element => {
    return (
        <div css={wrapper}>
            <div css={line} />
            <p css={text}>or</p>
            <div css={line} />
        </div>
    )
}

export default Separator
