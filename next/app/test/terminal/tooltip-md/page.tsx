import { GqlTerminalComponent } from "@/app/components/terminal2/GqlTerminalComponent";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTestTerminalTooltipMdPage {
    _test {
      appTestTerminalTooltipMdPage {
        ...GqlTerminalComponent
      }
    }
  }
`);

export default async function Page() {
  const data = await request("http://localhost:8080/query", queryDefinition);

  const fragment = data._test?.appTestTerminalTooltipMdPage;

  if (!fragment) {
    return <>no data</>;
  }

  return (
    <div style={{ height: "95svh" }}>
      <GqlTerminalComponent fragment={fragment} />
    </div>
  );
}
