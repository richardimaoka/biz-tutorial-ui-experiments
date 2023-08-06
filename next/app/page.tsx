import { graphql } from "@/libs/gql";
import { getClient } from "@/libs/apolloClient";
import RouterMounting from "./RouterMounting";
import { VisibleColumn } from "./components/column/VisibleColumn";

const queryDefinition = graphql(/* GraphQL */ `
  query PageQuery($tutorial: String!, $step: String, $openFilePath: String) {
    page(tutorial: $tutorial, step: $step) {
      ...VisibleColumn_Fragment
    }
  }
`);

interface PageParams {
  searchParams: {
    column?: string;
    step?: string;
    openFilePath?: string;
  };
}

export default async function Home({ searchParams }: PageParams) {
  const openFilePath = searchParams.openFilePath
    ? searchParams.openFilePath
    : "public/robots.txt";
  const step = "bf3aadbd-c876-4fd3-817b-3b0fc24b04f9";

  const { data } = await getClient().query({
    query: queryDefinition,
    variables: {
      tutorial: "sign-in-with-google",
      openFilePath: openFilePath,
      step: step,
      column: searchParams.column,
    },
  });

  return (
    <RouterMounting>
      <main>
        {data.page && (
          <VisibleColumn
            fragment={data.page}
            selectColumn={searchParams.column}
            openFilePath={openFilePath}
            step={step}
          />
        )}
      </main>
    </RouterMounting>
  );
}
