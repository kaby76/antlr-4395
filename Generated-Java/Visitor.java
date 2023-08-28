import org.antlr.v4.runtime.*;
import java.nio.charset.StandardCharsets;
import java.io.OutputStreamWriter;
import java.io.PrintWriter;
import java.io.PrintStream;

public class Visitor extends ExprBaseVisitor<Integer>
{
//    @Override
//    public Integer visitProg(ExprParser.ProgContext context)
//    {
//	System.out.println("VisitProg");
//	return this.visitChildren(context);
//    }

    @Override
    public Integer visitDecl(ExprParser.DeclContext context)
    {
	System.out.println("visitDecl");
	return this.visitChildren(context);
    }

    @Override
    public Integer visitExpr(ExprParser.ExprContext context)
    {
	System.out.println("visitExpr");
	return this.visitChildren(context);
    }
}
