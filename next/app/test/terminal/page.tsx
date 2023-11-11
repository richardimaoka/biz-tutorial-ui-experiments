import { GqlTerminalColumn } from "@/app/components/terminal/GqlTerminalColumn";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";
import Link from "next/link";

const queryDefinition = graphql(`
  query appTestTerminalPage($step: Int) {
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
  const stepNum = searchParams.step ? Number(searchParams.step) : 1;
  const nextStep = `${stepNum + 1}`;

  const variables = { step: stepNum };
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
      <Link href={`/test/terminal?step=${nextStep}`}>next</Link>
      <GqlTerminalColumn fragment={fragment} selectIndex={0} />
    </div>
  );
}
