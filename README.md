# go cabin

Similar to my goals in [ruby-cabin](https://github.com/jordansissel/ruby-cabin); I want good logging.

For now, while I experiment with git, here's what it 'stdout' output looks like:

    % ./simple 
    2012-04-13T05:35:59.318282Z: Hello world
    2012-04-13T05:35:59.318638Z: 42
    2012-04-13T05:35:59.318801Z: <Example Code(int)=42, Message(string)="The answer.">

You can log *any* object, including structs as you can see in the 3rd line above.
