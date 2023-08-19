import { graphql } from "@/libs/gql";
import { getClient } from "@/libs/apolloClient";
import { VisibleColumn } from "./components/column/VisibleColumn";
import { print } from "graphql";

const queryDefinition = graphql(/* GraphQL */ `
  query PageQuery($tutorial: String!, $step: String, $openFilePath: String) {
    page(tutorial: $tutorial, step: $step) {
      ...VisibleColumn_Fragment
      step
      focusColumn
      autoNextSeconds
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
  // console.log(print(queryDefinition));

  const openFilePath = searchParams.openFilePath;
  const { data } = await getClient().query({
    query: queryDefinition,
    variables: {
      // tutorial: "sign-in-with-google",
      tutorial: "live-server",
      openFilePath: openFilePath,
      step: searchParams.step,
      column: searchParams.column,
    },
  });

  let selectColumn: string | undefined;
  if (searchParams.column) {
    selectColumn = searchParams.column;
  } else if (data?.page?.focusColumn) {
    selectColumn = data.page.focusColumn;
  } else {
    selectColumn = undefined;
  }

  const step = searchParams.step ? searchParams.step : "_initial";

  return (
    <main>
      {data.page && (
        <VisibleColumn
          fragment={data.page}
          selectColumn={selectColumn}
          openFilePath={openFilePath}
          step={step}
          skipAnimation={searchParams.skipAnimation}
        />
      )}
    </main>
  );
}
