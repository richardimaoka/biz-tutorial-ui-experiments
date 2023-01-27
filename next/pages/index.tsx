import { useQuery } from "@apollo/client";
import { css } from "@emotion/react";
import Link from "next/link";
import { useRouter } from "next/router";
import { useEffect } from "react";
import { Header } from "../components/Header";
import { TerminalComponent } from "../components/terminal/TerminalComponent";
import { graphql } from "../libs/gql";

const Home2_Query = graphql(/* GraphQL */ `
  query Home2_Query($step: Int!) {
    terminal(step: $step) {
      ...TerminalComponent_Fragment
    }
  }
`);

export default function Home2() {
  const router = useRouter();
  const { step } = router.query;
  const stepInt = typeof step === "string" ? Math.trunc(Number(step)) : 0;

  const { loading, error, data } = useQuery(Home2_Query, {
    variables: { step: stepInt },
  });

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.code === "Space") {
        router.push(`./?step=${stepInt + 1}`);
      }
    };
    document.addEventListener("keydown", handleKeyDown);

    // Don't forget to clean up
    return function cleanup() {
      document.removeEventListener("keydown", handleKeyDown);
    };
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error : {error.message}</p>;

  return (
    data && (
      <>
        <Header />
        <main
          css={css`
            background-color: #333333;
          `}
          onKeyDown={(e) => e}
        >
          <div
            css={css`
              width: 680px;
              margin: 0 auto;
              background-color: white;
            `}
          >
            {data.terminal && <TerminalComponent fragment={data.terminal} />}
            <button type="button">
              <Link href={`./?step=${stepInt + 1}`}> next step</Link>
            </button>
            <input />
          </div>
        </main>
      </>
    )
  );
}
