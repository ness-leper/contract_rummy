import Head from "next/head";

import { api } from "~/utils/api";

export default function Home() {
  return (
    <>
      <Head>
        <title>Contract Rummy</title>
        <meta
          name="description"
          content="Keep track of scores using my rules."
        />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className="flex min-h-screen flex-col items-center bg-gradient-to-b from-[#2e026d] to-[#15162c]">
        <h1 className="text-5xl font-extrabold tracking-tight text-white sm:text-[5rem]">
          <span className="text-[hsl(280,100%,70%)]">Contract Rummy</span>
        </h1>
        <a href="/step1">
          <button className="mt-4 rounded bg-blue-500 px-4 py-2 font-bold text-white hover:bg-blue-700">
            Start Game
          </button>
        </a>
      </main>
    </>
  );
}
