import { useQuery } from "@apollo/client";
import { css } from "@emotion/react";
import { Header } from "../components/Header";
import { TerminalComponent } from "../components/terminal/TerminalComponent";
import { graphql } from "../libs/gql";

const Home2_Query = graphql(/* GraphQL */ `
  query Home2_Query {
    terminal {
      ...TerminalComponent_Fragment
    }
  }
`);

export default function Home2() {
  const { loading, error, data } = useQuery(Home2_Query);
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
        >
          <div
            css={css`
              width: 680px;
              margin: 0 auto;
              background-color: white;
            `}
          >
            {data.terminal && <TerminalComponent fragment={data.terminal} />}
          </div>
        </main>
      </>
    )
  );
}
