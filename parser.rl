// fasturl is a Go URL parser using a [Ragel](http://www.colm.net/open-source/ragel/) state-machine instead of regex, or the built in standard library `url.Parse`.
// 
package fasturl

import "fmt"

%%{
  machine url_parser;

  action mark { mark = fpc }
  action mark_port { port_mark = fpc }

  action save_port {
    if port_mark > host_mark{
      u.Port = data[port_mark:fpc]
    }
  }

  action save_scheme {
    u.Protocol = data[0:fpc-1]
  }

  action mark_host {
    host_mark = fpc;
  }

  action save_host {
    u.Host = data[host_mark:fpc]
  }

  action save_query {
    if u.Query == "" {
      u.Query = data[mark:fpc]
    }
  }

  action save_path {
    if u.Path == "" {
      u.Path = data[mark:fpc]
    }
  }

  action save_fragment {
    u.Fragment = data[mark:fpc]
  }

  pct_encoded = "%" xdigit xdigit;

  gen_delims  = ":" | "/" | "?" | "#" | "[" | "]" | "@";
  sub_delims  = "!" | "$" | "&" | "'" | "(" | ")" | "*" | "+" | "," | ";" | "=";

  reserved    = gen_delims | sub_delims;
  unreserved  = alpha | digit | "-" | "." | "_" | "~";

  # many clients don't encode these, e.g. curl, wget, ...
  delims      = "<" | ">" | "%" |  "#" | '"';
  unwise      = " " | "{" | "}" | "|" | "\\" | "^" | "[" | "]" | "`";

  pchar = unreserved | pct_encoded | sub_delims | ":" | "@" | delims | unwise;
  slash = "/" | "\\";
  path = (slash ( (pchar - ("?" | "#")) + ( slash (pchar - ("?" | "#"))* )* )? ) >mark %save_path;
  drivepath = ( (slash|(alpha ":" slash)) ( (pchar - ("?" | "#")) + ( slash (pchar - ("?" | "#"))* )* )? ) >mark %save_path;
  scheme = (alpha ( alpha | digit | "+" | "-" | "." )*);

  #simple ipv4 address
  dec_octet = digit{1,3};
  IPv4address = dec_octet "." dec_octet "." dec_octet "." dec_octet;

  IPvFuture  = "v" xdigit+ "." ( unreserved | sub_delims | ":" )+;

  # simple ipv6 address
  IPv6address = (":" | xdigit)+ IPv4address?;

  IP_literal = "[" ( IPv6address | IPvFuture  ) "]";

  reg_name = ( unreserved | pct_encoded | sub_delims )+;

  userinfo    = ( unreserved | pct_encoded | sub_delims | ":" | "@" )*;
  host        = IP_literal | IPv4address | reg_name;
  port        = (pchar - ("/" | "?" | "#")){1,5} ;
  authority   =  ( userinfo "@" )? (host >mark_host %save_host) ( ":" port >mark_port %save_port)?;

  fragment = ( pchar | "/" | "?" )* >mark %save_fragment;
  query = (pchar - "#")* >mark %save_query;

  full_ref = drivepath ( "?" query )? ( "#" fragment )?;
  relative_ref = path ( "?" query )? ( "#" fragment )?;
  absolute_hier_part = ("//")? authority? full_ref?;
  hier_part = ("//")? authority? relative_ref?;

  absolute_URI = ((scheme  ":") >mark %save_scheme)? absolute_hier_part;
  URI = absolute_URI | relative_ref;
  main := URI;

  write data;
}%%

// URL represents the different parts of a parsed URL
type URL struct {
  Protocol string
  Host     string
  Port     string
  Path     string
  Query    string
  Fragment string
}

// ParseURL parses a given URL and returns a `URL` representing the different parts
func ParseURL(data string) (*URL, error){
  mark, host_mark, port_mark, cs, p, pe, eof := 0, 0, 0, url_parser_en_main, 0, len(data), len(data)

  u := &URL{}
  
  %% write init;
  %% write exec;
  if cs < url_parser_first_final {
    return nil, fmt.Errorf("Failed to match URL")
  } else {
    return u, nil
  }
}

