interface iPrintRules {
  label: string
}

/**
 * @param {object} props
 * @param {string} props.label
 */
export function PrintRules(props:iPrintRules){
  return <div>{props.label}</div>;
}
