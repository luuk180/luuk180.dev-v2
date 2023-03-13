import type { AppProps } from 'next/app'
import 'bootstrap/dist/css/bootstrap.css'
import Head from "next/head";

export default function App({ Component, pageProps }: AppProps) {
  return <>
      <Head>
          <title>luuk180.dev</title>
      </Head>
      <Component {...pageProps} />
    </>
}
