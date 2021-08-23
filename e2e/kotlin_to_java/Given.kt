package earth.levi.dotenv

class FooClass: Activity {
    val combined = Env.barBar + Env.fooFoo
    val empty = Env.bazBaz
}