import { GqlTerminalColumn } from "@/app/components/terminal2/GqlTerminalColumn";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTestTerminalPage($step: String) {
    _test {
      appTestTerminalPage(step: $step) {
        ...GqlTerminalColumn
      }
    }
  }
`);

interface PageParams {
  searchParams: {
    step?: string;
  };
}

export default async function Page({ searchParams }: PageParams) {
  const variables = { step: searchParams.step };
  const data = await request(
    "http://localhost:8080/query",
    queryDefinition,
    variables
  );

  const fragment = data._test?.appTestTerminalPage;

  if (!fragment) {
    return <>no data</>;
  }

  return (
    <div style={{ height: "95svh" }}>
      <GqlTerminalColumn fragment={fragment} selectIndex={0} />
    </div>
  );
}
