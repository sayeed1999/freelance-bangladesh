'use client';
import React from 'react'

import { SessionProvider } from 'next-auth/react';

const SessionProviderWrapper = ({children}) => {
  return (
    <SessionProvider refetchInterval={4 * 60}>
      {children}
    </SessionProvider>
  )
}

export default SessionProviderWrapper;
