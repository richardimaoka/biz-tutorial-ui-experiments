import { graphql } from "@/libs/gql";
import { getClient } from "@/libs/apolloClient";
import RouterMounting from "./RouterMounting";
import { VisibleColumn } from "./components/column/VisibleColumn";
import { Navigation } from "./components/navigation/Navigation";

import { Noto_Sans_JP } from "next/font/google";

const notojp = Noto_Sans_JP({
  weight: "400",
  preload: false,
  display: "swap", // フォントの表示方法を指定します
});

const queryDefinition = graphql(/* GraphQL */ `
  query PageQuery($tutorial: String!, $step: String, $openFilePath: String) {
    page(tutorial: $tutorial, step: $step) {
      ...VisibleColumn_Fragment
      ...Navigation_Fragment
      step
    }
  }
`);

interface PageParams {
  searchParams: {
    column?: string;
    step?: string;
    openFilePath?: string;
    skipAnimation?: boolean;
  };
}

export default async function Home({ searchParams }: PageParams) {
  const openFilePath = searchParams.openFilePath
    ? searchParams.openFilePath
    : "package.json";

  const { data } = await getClient().query({
    query: queryDefinition,
    variables: {
      tutorial: "sign-in-with-google",
      openFilePath: openFilePath,
      step: searchParams.step,
      column: searchParams.column,
    },
  });

  return (
    <RouterMounting>
      <main className={notojp.className}>
        {data.page && (
          <>
            <VisibleColumn
              fragment={data.page}
              selectColumn={searchParams.column}
              openFilePath={openFilePath}
              step={searchParams.step}
              skipAnimation={searchParams.skipAnimation}
            />
            <Navigation fragment={data.page} />
          </>
        )}
      </main>
    </RouterMounting>
  );
}
