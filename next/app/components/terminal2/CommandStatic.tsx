interface Props {
  command: string;
}

export function CommandStatic(props: Props) {
  return <code>{props.command}</code>;
}
